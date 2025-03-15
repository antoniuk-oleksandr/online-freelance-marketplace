import {ResponseErrorEnum} from "@/types/ResponseErrorEnum";

export type ErrorStore = {
    error: ResponseErrorEnum | undefined;
    shown: boolean;
}
