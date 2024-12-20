import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";

export type ErrorStore = {
    error: ResponseErrorEnum | undefined;
    shown: boolean;
}