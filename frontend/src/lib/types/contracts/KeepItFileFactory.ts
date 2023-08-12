/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BytesLike,
  FunctionFragment,
  Result,
  Interface,
  EventFragment,
  AddressLike,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedLogDescription,
  TypedListener,
  TypedContractMethod,
} from "./common";

export interface KeepItFileFactoryInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "changeOwner"
      | "createFile"
      | "getAttestationServiceAddres"
      | "getFiles"
      | "getOwner"
      | "setAttestationServiceAddress"
  ): FunctionFragment;

  getEvent(nameOrSignatureOrTopic: "FileCreated" | "OwnerSet"): EventFragment;

  encodeFunctionData(
    functionFragment: "changeOwner",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "createFile",
    values: [string, string, string, string[], string[]]
  ): string;
  encodeFunctionData(
    functionFragment: "getAttestationServiceAddres",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getFiles",
    values: [AddressLike]
  ): string;
  encodeFunctionData(functionFragment: "getOwner", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "setAttestationServiceAddress",
    values: [AddressLike]
  ): string;

  decodeFunctionResult(
    functionFragment: "changeOwner",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "createFile", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getAttestationServiceAddres",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "getFiles", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "getOwner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "setAttestationServiceAddress",
    data: BytesLike
  ): Result;
}

export namespace FileCreatedEvent {
  export type InputTuple = [fileAddress: AddressLike];
  export type OutputTuple = [fileAddress: string];
  export interface OutputObject {
    fileAddress: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace OwnerSetEvent {
  export type InputTuple = [oldOwner: AddressLike, newOwner: AddressLike];
  export type OutputTuple = [oldOwner: string, newOwner: string];
  export interface OutputObject {
    oldOwner: string;
    newOwner: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface KeepItFileFactory extends BaseContract {
  connect(runner?: ContractRunner | null): KeepItFileFactory;
  waitForDeployment(): Promise<this>;

  interface: KeepItFileFactoryInterface;

  queryFilter<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;
  queryFilter<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;

  on<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  on<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  once<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  once<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  listeners<TCEvent extends TypedContractEvent>(
    event: TCEvent
  ): Promise<Array<TypedListener<TCEvent>>>;
  listeners(eventName?: string): Promise<Array<Listener>>;
  removeAllListeners<TCEvent extends TypedContractEvent>(
    event?: TCEvent
  ): Promise<this>;

  changeOwner: TypedContractMethod<
    [newOwner: AddressLike],
    [void],
    "nonpayable"
  >;

  createFile: TypedContractMethod<
    [
      _fileType: string,
      _fileName: string,
      _fileCheckSum: string,
      _transactionList: string[],
      _chainIdList: string[]
    ],
    [void],
    "nonpayable"
  >;

  getAttestationServiceAddres: TypedContractMethod<[], [string], "view">;

  getFiles: TypedContractMethod<[_owner: AddressLike], [string[]], "view">;

  getOwner: TypedContractMethod<[], [string], "view">;

  setAttestationServiceAddress: TypedContractMethod<
    [_aSA: AddressLike],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "changeOwner"
  ): TypedContractMethod<[newOwner: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "createFile"
  ): TypedContractMethod<
    [
      _fileType: string,
      _fileName: string,
      _fileCheckSum: string,
      _transactionList: string[],
      _chainIdList: string[]
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "getAttestationServiceAddres"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "getFiles"
  ): TypedContractMethod<[_owner: AddressLike], [string[]], "view">;
  getFunction(
    nameOrSignature: "getOwner"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "setAttestationServiceAddress"
  ): TypedContractMethod<[_aSA: AddressLike], [void], "nonpayable">;

  getEvent(
    key: "FileCreated"
  ): TypedContractEvent<
    FileCreatedEvent.InputTuple,
    FileCreatedEvent.OutputTuple,
    FileCreatedEvent.OutputObject
  >;
  getEvent(
    key: "OwnerSet"
  ): TypedContractEvent<
    OwnerSetEvent.InputTuple,
    OwnerSetEvent.OutputTuple,
    OwnerSetEvent.OutputObject
  >;

  filters: {
    "FileCreated(address)": TypedContractEvent<
      FileCreatedEvent.InputTuple,
      FileCreatedEvent.OutputTuple,
      FileCreatedEvent.OutputObject
    >;
    FileCreated: TypedContractEvent<
      FileCreatedEvent.InputTuple,
      FileCreatedEvent.OutputTuple,
      FileCreatedEvent.OutputObject
    >;

    "OwnerSet(address,address)": TypedContractEvent<
      OwnerSetEvent.InputTuple,
      OwnerSetEvent.OutputTuple,
      OwnerSetEvent.OutputObject
    >;
    OwnerSet: TypedContractEvent<
      OwnerSetEvent.InputTuple,
      OwnerSetEvent.OutputTuple,
      OwnerSetEvent.OutputObject
    >;
  };
}
