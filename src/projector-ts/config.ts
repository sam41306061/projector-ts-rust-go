import { config } from 'process';
import {Ops} from '../../opts';
import path from 'path';
export enum Operation {
    Print,
    Add,
    Remove
}

export type Config = {
    args: string[],
    operation: Operation
    config: string
    pwd: string
}

function getPwd(opts: Ops): string {
    if(opts.pwd) {
        //@ts-ignore
        return opts.pwd
    }
    return process.cwd()
}

function getConfigPath(ops: Ops): string {
    if(ops.config) {
        return ops.config
    }
    const home = process.env["HOME"]
    const loc = process.env["XDG_CONFIG_HOME"] || home; 
    if(!loc) {
        throw new Error("unable to determine config location");
    } 
    if(loc === home) {
        return path.join(loc, ".projector.json")
    }
    return path.join(loc, "projector", "projector.json")
}

function getOperation(opts:Ops): Operation {
    if(!opts.args || opts.args.length === 0) {
        return Operation.Print
    }
    if(opts[0] === "add") {
        return Operation.Add
    }
    if(opts[0] === "remove") {
        return Operation.Remove
    }
    return Operation.Print;
}

function getArgs(opts: Ops): string[] {
    if(!opts.args || opts.args.length === 0) {
        return [];
    }
    const operation = getOperation(opts);
    if(operation === Operation.Print) {
        if(opts.args.length > 1){
            throw new Error(`expected 0 or 1 args but got:${opts.args.length}`);
        }
        return opts.args;
    }
    if(operation === Operation.Add) {
        if(opts.args.length !== 3){
            throw new Error(`expected 2  but got:${opts.args.length - 1}`);
        }
        return opts.args.slice(1);
    }
    if(opts.args.length !== 2) {
        throw new Error(`expected 1  but got:${opts.args.length - 1}`);
    }
    return opts.args.slice(1);
}
// function to use 
export default function getConfig(opts: Ops): Config {
    return {
        pwd: getPwd(opts),
        config: getConfigPath(opts),
        args: getArgs(opts),
        operation: getOperation(opts)
    };
}