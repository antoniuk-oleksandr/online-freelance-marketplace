import {SignUpData} from "@/types/SignData.ts";
import {getHost} from "@/utils/utils.ts";
import axios from "axios";

export const postSignUpRequest = async (signUpData: SignUpData) => {
    const host = getHost();
    const url = `http://${host}/api/v1/auth/sign-up`;

    try {
        const response = await axios.post(url, signUpData);
        return {
            status: response.status,
            data: response.data
        }
    }
    catch (e) {
        return {
            status: (e as any).response.status,
            data: (e as any).response.data
        }
    }
}