/* eslint-disable */
import { BaseAccount } from '../cosmos/auth/v1beta1/auth'
import { Grant } from '../dclauth/grant'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth'

export interface Account {
  baseAccount: BaseAccount | undefined
  /**
   * NOTE. we do not user AccountRoles casting here to preserve repeated form
   *       so protobuf takes care about repeated items in generated code,
   *       (but that might be not the final solution)
   */
  roles: string[]
  approvals: Grant[]
  vendorID: number
}

const baseAccount: object = { roles: '', vendorID: 0 }

export const Account = {
  encode(message: Account, writer: Writer = Writer.create()): Writer {
    if (message.baseAccount !== undefined) {
      BaseAccount.encode(message.baseAccount, writer.uint32(10).fork()).ldelim()
    }
    for (const v of message.roles) {
      writer.uint32(18).string(v!)
    }
    for (const v of message.approvals) {
      Grant.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    if (message.vendorID !== 0) {
      writer.uint32(32).int32(message.vendorID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Account {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseAccount } as Account
    message.roles = []
    message.approvals = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.baseAccount = BaseAccount.decode(reader, reader.uint32())
          break
        case 2:
          message.roles.push(reader.string())
          break
        case 3:
          message.approvals.push(Grant.decode(reader, reader.uint32()))
          break
        case 4:
          message.vendorID = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Account {
    const message = { ...baseAccount } as Account
    message.roles = []
    message.approvals = []
    if (object.baseAccount !== undefined && object.baseAccount !== null) {
      message.baseAccount = BaseAccount.fromJSON(object.baseAccount)
    } else {
      message.baseAccount = undefined
    }
    if (object.roles !== undefined && object.roles !== null) {
      for (const e of object.roles) {
        message.roles.push(String(e))
      }
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(Grant.fromJSON(e))
      }
    }
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = Number(object.vendorID)
    } else {
      message.vendorID = 0
    }
    return message
  },

  toJSON(message: Account): unknown {
    const obj: any = {}
    message.baseAccount !== undefined && (obj.baseAccount = message.baseAccount ? BaseAccount.toJSON(message.baseAccount) : undefined)
    if (message.roles) {
      obj.roles = message.roles.map((e) => e)
    } else {
      obj.roles = []
    }
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => (e ? Grant.toJSON(e) : undefined))
    } else {
      obj.approvals = []
    }
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    return obj
  },

  fromPartial(object: DeepPartial<Account>): Account {
    const message = { ...baseAccount } as Account
    message.roles = []
    message.approvals = []
    if (object.baseAccount !== undefined && object.baseAccount !== null) {
      message.baseAccount = BaseAccount.fromPartial(object.baseAccount)
    } else {
      message.baseAccount = undefined
    }
    if (object.roles !== undefined && object.roles !== null) {
      for (const e of object.roles) {
        message.roles.push(e)
      }
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(Grant.fromPartial(e))
      }
    }
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = object.vendorID
    } else {
      message.vendorID = 0
    }
    return message
  }
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
