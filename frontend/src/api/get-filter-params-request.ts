import {getHost} from "@/utils/utils.ts";
import axios from "axios";
import {Skill} from "@/types/Skill.ts";
import {Category} from "@/types/Category.ts";
import {Language} from "@/types/Language.ts";
import {GetFilterParamsRequestResponse} from "@/types/GetFilterParamsRequestResponse";
import {ResponseError} from "@/types/ResponseErrorEnum.ts";

export const getFilterParamsRequest = async (

): Promise<GetFilterParamsRequestResponse> => {
    const host = getHost();
    const url = `http://${host}/api/v1/filter-params/get-all`;

    try {
        const response = await axios.get(url);
        return {
            data: response.data,
            status: response.status,
        }
    } catch (e) {
        return {
            status: (e as any).response.status,
            error: (e as any).response.data.message
        }
    }
}