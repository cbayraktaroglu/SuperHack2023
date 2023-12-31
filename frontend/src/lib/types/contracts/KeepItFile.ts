/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
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

export declare namespace KeepItFile {
  export type FilePartsStruct = { chainID: string; transactionHash: string };

  export type FilePartsStructOutput = [
    chainID: string,
    transactionHash: string
  ] & { chainID: string; transactionHash: string };
}

export interface KeepItFileInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "checkVerification"
      | "getFileCheckSum"
      | "getFileName"
      | "getFileType"
      | "getNumberofTransactions"
      | "getOrgVerification"
      | "getOwner"
      | "getTransactionAtIndex"
      | "getVerification"
      | "orgVerify"
      | "transferOwnership"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic: "OwnershipTransferred" | "VerificationResult"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "checkVerification",
    values: [
      string,
      string,
      AddressLike,
      BigNumberish,
      BigNumberish,
      BigNumberish[]
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "getFileCheckSum",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getFileName",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getFileType",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getNumberofTransactions",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getOrgVerification",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "getOwner", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "getTransactionAtIndex",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "getVerification",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "orgVerify",
    values: [BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "transferOwnership",
    values: [AddressLike]
  ): string;

  decodeFunctionResult(
    functionFragment: "checkVerification",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getFileCheckSum",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getFileName",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getFileType",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getNumberofTransactions",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getOrgVerification",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "getOwner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getTransactionAtIndex",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getVerification",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "orgVerify", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "transferOwnership",
    data: BytesLike
  ): Result;
}

export namespace OwnershipTransferredEvent {
  export type InputTuple = [
    fileAddress: AddressLike,
    previousOwner: AddressLike,
    newOwner: AddressLike
  ];
  export type OutputTuple = [
    fileAddress: string,
    previousOwner: string,
    newOwner: string
  ];
  export interface OutputObject {
    fileAddress: string;
    previousOwner: string;
    newOwner: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace VerificationResultEvent {
  export type InputTuple = [owner: AddressLike, result: boolean];
  export type OutputTuple = [owner: string, result: boolean];
  export interface OutputObject {
    owner: string;
    result: boolean;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface KeepItFile extends BaseContract {
  connect(runner?: ContractRunner | null): KeepItFile;
  waitForDeployment(): Promise<this>;

  interface: KeepItFileInterface;

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

  checkVerification: TypedContractMethod<
    [
      _appId: string,
      _actionId: string,
      signal: AddressLike,
      root: BigNumberish,
      nullifierHash: BigNumberish,
      proof: BigNumberish[]
    ],
    [void],
    "nonpayable"
  >;

  getFileCheckSum: TypedContractMethod<[], [string], "view">;

  getFileName: TypedContractMethod<[], [string], "view">;

  getFileType: TypedContractMethod<[], [string], "view">;

  getNumberofTransactions: TypedContractMethod<[], [bigint], "view">;

  getOrgVerification: TypedContractMethod<[], [boolean], "view">;

  getOwner: TypedContractMethod<[], [string], "view">;

  getTransactionAtIndex: TypedContractMethod<
    [index: BigNumberish],
    [KeepItFile.FilePartsStructOutput],
    "view"
  >;

  getVerification: TypedContractMethod<[], [boolean], "view">;

  orgVerify: TypedContractMethod<[_uid: BytesLike], [void], "nonpayable">;

  transferOwnership: TypedContractMethod<
    [newOwner: AddressLike],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "checkVerification"
  ): TypedContractMethod<
    [
      _appId: string,
      _actionId: string,
      signal: AddressLike,
      root: BigNumberish,
      nullifierHash: BigNumberish,
      proof: BigNumberish[]
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "getFileCheckSum"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "getFileName"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "getFileType"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "getNumberofTransactions"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "getOrgVerification"
  ): TypedContractMethod<[], [boolean], "view">;
  getFunction(
    nameOrSignature: "getOwner"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "getTransactionAtIndex"
  ): TypedContractMethod<
    [index: BigNumberish],
    [KeepItFile.FilePartsStructOutput],
    "view"
  >;
  getFunction(
    nameOrSignature: "getVerification"
  ): TypedContractMethod<[], [boolean], "view">;
  getFunction(
    nameOrSignature: "orgVerify"
  ): TypedContractMethod<[_uid: BytesLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "transferOwnership"
  ): TypedContractMethod<[newOwner: AddressLike], [void], "nonpayable">;

  getEvent(
    key: "OwnershipTransferred"
  ): TypedContractEvent<
    OwnershipTransferredEvent.InputTuple,
    OwnershipTransferredEvent.OutputTuple,
    OwnershipTransferredEvent.OutputObject
  >;
  getEvent(
    key: "VerificationResult"
  ): TypedContractEvent<
    VerificationResultEvent.InputTuple,
    VerificationResultEvent.OutputTuple,
    VerificationResultEvent.OutputObject
  >;

  filters: {
    "OwnershipTransferred(address,address,address)": TypedContractEvent<
      OwnershipTransferredEvent.InputTuple,
      OwnershipTransferredEvent.OutputTuple,
      OwnershipTransferredEvent.OutputObject
    >;
    OwnershipTransferred: TypedContractEvent<
      OwnershipTransferredEvent.InputTuple,
      OwnershipTransferredEvent.OutputTuple,
      OwnershipTransferredEvent.OutputObject
    >;

    "VerificationResult(address,bool)": TypedContractEvent<
      VerificationResultEvent.InputTuple,
      VerificationResultEvent.OutputTuple,
      VerificationResultEvent.OutputObject
    >;
    VerificationResult: TypedContractEvent<
      VerificationResultEvent.InputTuple,
      VerificationResultEvent.OutputTuple,
      VerificationResultEvent.OutputObject
    >;
  };
}
