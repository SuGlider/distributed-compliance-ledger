/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Plan } from '../cosmos/upgrade/v1beta1/upgrade'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade'

export interface MsgProposeUpgrade {
  creator: string
  plan: Plan | undefined
}

export interface MsgProposeUpgradeResponse {}

export interface MsgApproveUpgrade {
  creator: string
  name: string
}

export interface MsgApproveUpgradeResponse {}

const baseMsgProposeUpgrade: object = { creator: '' }

export const MsgProposeUpgrade = {
  encode(message: MsgProposeUpgrade, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.plan !== undefined) {
      Plan.encode(message.plan, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeUpgrade {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeUpgrade } as MsgProposeUpgrade
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.plan = Plan.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgProposeUpgrade {
    const message = { ...baseMsgProposeUpgrade } as MsgProposeUpgrade
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.plan !== undefined && object.plan !== null) {
      message.plan = Plan.fromJSON(object.plan)
    } else {
      message.plan = undefined
    }
    return message
  },

  toJSON(message: MsgProposeUpgrade): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.plan !== undefined && (obj.plan = message.plan ? Plan.toJSON(message.plan) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<MsgProposeUpgrade>): MsgProposeUpgrade {
    const message = { ...baseMsgProposeUpgrade } as MsgProposeUpgrade
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.plan !== undefined && object.plan !== null) {
      message.plan = Plan.fromPartial(object.plan)
    } else {
      message.plan = undefined
    }
    return message
  }
}

const baseMsgProposeUpgradeResponse: object = {}

export const MsgProposeUpgradeResponse = {
  encode(_: MsgProposeUpgradeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeUpgradeResponse } as MsgProposeUpgradeResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgProposeUpgradeResponse {
    const message = { ...baseMsgProposeUpgradeResponse } as MsgProposeUpgradeResponse
    return message
  },

  toJSON(_: MsgProposeUpgradeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgProposeUpgradeResponse>): MsgProposeUpgradeResponse {
    const message = { ...baseMsgProposeUpgradeResponse } as MsgProposeUpgradeResponse
    return message
  }
}

const baseMsgApproveUpgrade: object = { creator: '', name: '' }

export const MsgApproveUpgrade = {
  encode(message: MsgApproveUpgrade, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.name !== '') {
      writer.uint32(18).string(message.name)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveUpgrade {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveUpgrade } as MsgApproveUpgrade
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.name = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgApproveUpgrade {
    const message = { ...baseMsgApproveUpgrade } as MsgApproveUpgrade
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    return message
  },

  toJSON(message: MsgApproveUpgrade): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.name !== undefined && (obj.name = message.name)
    return obj
  },

  fromPartial(object: DeepPartial<MsgApproveUpgrade>): MsgApproveUpgrade {
    const message = { ...baseMsgApproveUpgrade } as MsgApproveUpgrade
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    return message
  }
}

const baseMsgApproveUpgradeResponse: object = {}

export const MsgApproveUpgradeResponse = {
  encode(_: MsgApproveUpgradeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveUpgradeResponse } as MsgApproveUpgradeResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgApproveUpgradeResponse {
    const message = { ...baseMsgApproveUpgradeResponse } as MsgApproveUpgradeResponse
    return message
  },

  toJSON(_: MsgApproveUpgradeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgApproveUpgradeResponse>): MsgApproveUpgradeResponse {
    const message = { ...baseMsgApproveUpgradeResponse } as MsgApproveUpgradeResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  ProposeUpgrade(request: MsgProposeUpgrade): Promise<MsgProposeUpgradeResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ApproveUpgrade(request: MsgApproveUpgrade): Promise<MsgApproveUpgradeResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  ProposeUpgrade(request: MsgProposeUpgrade): Promise<MsgProposeUpgradeResponse> {
    const data = MsgProposeUpgrade.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Msg', 'ProposeUpgrade', data)
    return promise.then((data) => MsgProposeUpgradeResponse.decode(new Reader(data)))
  }

  ApproveUpgrade(request: MsgApproveUpgrade): Promise<MsgApproveUpgradeResponse> {
    const data = MsgApproveUpgrade.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Msg', 'ApproveUpgrade', data)
    return promise.then((data) => MsgApproveUpgradeResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
