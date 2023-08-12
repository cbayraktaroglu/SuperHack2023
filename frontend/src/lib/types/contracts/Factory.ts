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

export interface FactoryInterface extends Interface {
  getFunction(nameOrSignature: "createFile" | "getFiles"): FunctionFragment;

  getEvent(nameOrSignatureOrTopic: "FileCreated"): EventFragment;

  encodeFunctionData(
    functionFragment: "createFile",
    values: [string, string, string, string[], string[]]
  ): string;
  encodeFunctionData(
    functionFragment: "getFiles",
    values: [AddressLike]
  ): string;

  decodeFunctionResult(functionFragment: "createFile", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "getFiles", data: BytesLike): Result;
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

export interface Factory extends BaseContract {
  connect(runner?: ContractRunner | null): Factory;
  waitForDeployment(): Promise<this>;

  interface: FactoryInterface;

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

  getFiles: TypedContractMethod<[_owner: AddressLike], [string[]], "view">;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

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
    nameOrSignature: "getFiles"
  ): TypedContractMethod<[_owner: AddressLike], [string[]], "view">;

  getEvent(
    key: "FileCreated"
  ): TypedContractEvent<
    FileCreatedEvent.InputTuple,
    FileCreatedEvent.OutputTuple,
    FileCreatedEvent.OutputObject
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
  };
}
