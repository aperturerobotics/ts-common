// @generated by protoc-gen-es-lite unknown with parameter "target=ts,ts_nocheck=false"
// @generated from file github.com/aperturerobotics/common/example/other/other.proto (package example.other, syntax proto3)
/* eslint-disable */

import type { MessageType, PartialFieldInfo } from '@aptre/protobuf-es-lite'
import { createMessageType, ScalarType } from '@aptre/protobuf-es-lite'

export const protobufPackage = 'example.other'

/**
 * OtherMsg is a different message from ExampleMsg.
 *
 * @generated from message example.other.OtherMsg
 */
export interface OtherMsg {
  /**
   * FooField is an example field.
   *
   * @generated from field: uint32 foo_field = 1;
   */
  fooField?: number
}

// OtherMsg contains the message type declaration for OtherMsg.
export const OtherMsg: MessageType<OtherMsg> = createMessageType({
  typeName: 'example.other.OtherMsg',
  fields: [
    { no: 1, name: 'foo_field', kind: 'scalar', T: ScalarType.UINT32 },
  ] as readonly PartialFieldInfo[],
  packedByDefault: true,
})
