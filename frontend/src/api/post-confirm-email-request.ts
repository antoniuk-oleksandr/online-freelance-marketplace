import {getHost} from "@/utils/utils.ts";
import axios from "axios";
import {PostConfirmEmailRequestResponse} from "@/types/PostConfirmEmailRequestResponse.ts";

export const postConfirmEmailRequest =
    async (token: string): Promise<PostConfirmEmailRequestResponse> => {
        const host = getHost();
        const url = `http://${host}/api/v1/auth/confirm-email`;

        try {
            const response = await axios.post(url, {}, {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            });
            return {
                data: response.data,
                status: response.status,
            }
        } catch (e) {
            return {
                data: (e as any).response.data,
                status: (e as any).response.status as number,
            }
        }
    }