import { AuthInfo, CommandResult, ObjectRef, ResultObject, ShortName, ValidateResult } from "./models";
import _ from "lodash";
import React from "react";
import { Box, Typography } from "@mui/material";
import Tooltip from "@mui/material/Tooltip";
import "./staticbuild.d.ts"
import { loadScript } from "./loadscript";
import { GitRef } from "./models-static";

console.log(window.location)

export let rootPath = window.location.pathname
if (rootPath.endsWith("/")) {
    rootPath = rootPath.substring(0, rootPath.length-1)
}
if (rootPath.endsWith("index.html")) {
    rootPath = rootPath.substring(0, rootPath.length-"index.html".length-1)
}
const staticPath = rootPath + "/staticdata"

console.log("rootPath=" + rootPath)
console.log("staticPath=" + staticPath)

export enum ObjectType {
    Rendered = "rendered",
    Remote = "remote",
    Applied = "applied",
}

export interface User {
    username: string
    isAdmin: boolean
}

export interface Api {
    getAuthInfo(): Promise<AuthInfo>
    getUser(): Promise<User>
    getShortNames(): Promise<ShortName[]>
    listenUpdates(filterProject: string | undefined, filterSubDir: string | undefined, handle: (msg: any) => void): Promise<() => void>
    getCommandResult(resultId: string): Promise<CommandResult>
    getCommandResultObject(resultId: string, ref: ObjectRef, objectType: string): Promise<any>
    getValidateResult(resultId: string): Promise<ValidateResult>
    validateNow(cluster: string, name: string, namespace: string): Promise<Response>
    reconcileNow(cluster: string, name: string, namespace: string): Promise<Response>
    deployNow(cluster: string, name: string, namespace: string): Promise<Response>
    setSuspended(cluster: string, name: string, namespace: string, suspend: boolean): Promise<Response>
}

export async function checkStaticBuild() {
    const p = loadScript(staticPath + "/summaries.js")
    try {
        await p
        return true
    } catch (error) {
        return false
    }
}

export class RealApi implements Api {
    onUnauthorized?: () => void;

    constructor(onUnauthorized?: () => void) {
        this.onUnauthorized = onUnauthorized
    }

    handleErrors(response: Response) {
        if (!response.ok) {
            if (response.status === 401) {
                if (this.onUnauthorized) {
                    this.onUnauthorized()
                }
            }
            throw Error(response.statusText)
        }
    }

    async doGet(path: string, params?: URLSearchParams, abort?: AbortSignal) {
        let url = rootPath + path
        if (params) {
            url += "?" + params.toString()
        }
        const resp = await fetch(url, {
            method: "GET",
            signal: abort ? abort : null,
        })
        this.handleErrors(resp)
        return resp.json()
    }

    async doPost(path: string, body: any) {
        let url = rootPath + path
        const headers = {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
        const resp = await fetch(url, {
            method: "POST",
            body: JSON.stringify(body),
            headers: headers,
        })
        this.handleErrors(resp)
        return resp
    }

    async getAuthInfo(): Promise<AuthInfo> {
        return this.doGet("/auth/info")
    }

    async getUser(): Promise<User> {
        return this.doGet("/auth/user")
    }

    async getShortNames(): Promise<ShortName[]> {
        return this.doGet("/api/getShortNames")
    }

    async listenUpdates(filterProject: string | undefined, filterSubDir: string | undefined, handle: (msg: any) => void): Promise<() => void> {
        const params = new URLSearchParams()
        if (filterProject) {
            params.set("filterProject", filterProject)
        }
        if (filterSubDir) {
            params.set("filterSubDir", filterSubDir)
        }

        let seq = 0
        const abort = new AbortController()

        const doGetEvents = async () => {
            if (abort.signal.aborted) {
                return
            }

            params.set("seq", seq + "")
            let resp: any
            try {
                resp = await this.doGet("/api/events", params, abort.signal)
            } catch (error) {
                console.log("events error", error)
                seq = 0
                await new Promise(r => setTimeout(r, 5000));
                doGetEvents()
                return
            }

            seq = resp.nextSeq
            const events = resp.events

            events.forEach((e: any) => {
                handle(e)
            })

            doGetEvents()
        }

        doGetEvents()

        return () => {
            console.log("events cancel")
            abort.abort()
        }
    }

    async getCommandResult(resultId: string) {
        const params = new URLSearchParams()
        params.set("resultId", resultId)
        const json = await this.doGet("/api/getCommandResult", params)
        return new CommandResult(json)
    }

    async getCommandResultObject(resultId: string, ref: ObjectRef, objectType: string) {
        const params = new URLSearchParams()
        params.set("resultId", resultId)
        params.set("objectType", objectType)
        buildRefParams(ref, params)
        return await this.doGet("/api/getCommandResultObject", params)
    }

    async getValidateResult(resultId: string) {
        const params = new URLSearchParams()
        params.set("resultId", resultId)
        const json = await this.doGet("/api/getValidateResult", params)
        return new ValidateResult(json)
    }

    async validateNow(cluster: string, name: string, namespace: string) {
        return this.doPost("/api/validateNow", {
            "cluster": cluster,
            "name": name,
            "namespace": namespace,
        })
    }

    async deployNow(cluster: string, name: string, namespace: string): Promise<Response> {
        return this.doPost("/api/deployNow", {
            "cluster": cluster,
            "name": name,
            "namespace": namespace,
        })
    }

    async reconcileNow(cluster: string, name: string, namespace: string): Promise<Response> {
        return this.doPost("/api/reconcileNow", {
            "cluster": cluster,
            "name": name,
            "namespace": namespace,
        })
    }

    async setSuspended(cluster: string, name: string, namespace: string, suspend: boolean): Promise<Response> {
        return this.doPost("/api/setSuspended", {
            "cluster": cluster,
            "name": name,
            "namespace": namespace,
            "suspend": suspend,
        })
    }
}

export class StaticApi implements Api {
    async getAuthInfo(): Promise<AuthInfo> {
        const info = new AuthInfo()
        info.authEnabled = false
        info.adminEnabled = false
        return info
    }

    async getUser(): Promise<User> {
        return {
            "username": "no-user",
            "isAdmin": true,
        }
    }

    async getShortNames(): Promise<ShortName[]> {
        await loadScript(staticPath + "/shortnames.js")
        return staticShortNames
    }

    async listenUpdates(filterProject: string | undefined, filterSubDir: string | undefined, handle: (msg: any) => void): Promise<() => void> {
        await loadScript(staticPath + "/summaries.js")

        staticSummaries.forEach(rs => {
            if (filterProject && filterProject !== rs.project.normalizedGitUrl) {
                return
            }
            if (filterSubDir && filterSubDir !== rs.project.subDir) {
                return
            }
            handle({
                "type": "update_summary",
                "summary": rs,
            })
        })
        return () => {
        }
    }

    async getCommandResult(resultId: string): Promise<CommandResult> {
        await loadScript(staticPath + `/result-${resultId}.js`)
        return staticResults.get(resultId)
    }

    async getCommandResultObject(resultId: string, ref: ObjectRef, objectType: string): Promise<any> {
        const result = await this.getCommandResult(resultId)
        const object = result.objects?.find(x => _.isEqual(x.ref, ref))
        if (!object) {
            throw new Error("object not found")
        }
        switch (objectType) {
            case ObjectType.Rendered:
                return object.rendered
            case ObjectType.Remote:
                return object.remote
            case ObjectType.Applied:
                return object.applied
            default:
                throw new Error("unknown object type " + objectType)
        }
    }

    async getValidateResult(resultId: string): Promise<ValidateResult> {
        throw new Error("not implemented")
    }

    validateNow(cluster: string, name: string, namespace: string): Promise<Response> {
        throw new Error("not implemented")
    }

    reconcileNow(cluster: string, name: string, namespace: string): Promise<Response> {
        throw new Error("not implemented")
    }

    deployNow(cluster: string, name: string, namespace: string): Promise<Response> {
        throw new Error("not implemented")
    }

    setSuspended(cluster: string, name: string, namespace: string, suspend: boolean): Promise<Response> {
        throw new Error("not implemented")
    }
}

function buildRefParams(ref: ObjectRef, params: URLSearchParams) {
    params.set("kind", ref.kind)
    params.set("name", ref.name)
    if (ref.group) {
        params.set("group", ref.group)
    }
    if (ref.version) {
        params.set("version", ref.version)
    }
    if (ref.namespace) {
        params.set("namespace", ref.namespace)
    }
}

export function buildRefString(ref: ObjectRef): string {
    if (ref.namespace) {
        return `${ref.namespace}/${ref.kind}/${ref.name}`
    } else {
        if (ref.name) {
            return `${ref.kind}/${ref.name}`
        } else {
            return ref.kind
        }
    }
}

export function buildRefKindElement(ref: ObjectRef, element?: React.ReactElement): React.ReactElement {
    const tooltip = <Box zIndex={1000}>
        <Typography>ApiVersion: {[ref.group, ref.version].filter(x => x).join("/")}</Typography><br/>
        <Typography>Kind: {ref.kind}</Typography>
    </Box>
    return <Tooltip title={tooltip}>
        {element ? element : <Typography>{ref.kind}</Typography>}
    </Tooltip>
}

export function buildObjectRefFromObject(obj: any): ObjectRef {
    const apiVersion: string = obj.apiVersion
    const s = apiVersion.split("/", 2)
    let ref = new ObjectRef()
    if (s.length === 1) {
        ref.version = s[0]
    } else {
        ref.group = s[0]
        ref.version = s[1]
    }
    ref.kind = obj.kind
    ref.namespace = obj.metadata.namespace
    ref.name = obj.metadata.name
    return ref
}

export function buildGitRefString(ref?: GitRef) {
    if (!ref) {
        return "HEAD"
    }
    if (ref.branch) {
        return ref.branch
    } else if (ref.tag) {
        return ref.tag
    } else {
        return "<unknown>"
    }
}

export function findObjectByRef(l: ResultObject[] | undefined, ref: ObjectRef, filter?: (o: ResultObject) => boolean): ResultObject | undefined {
    return l?.find(x => {
        if (filter && !filter(x)) {
            return false
        }
        return _.isEqual(x.ref, ref)
    })
}
