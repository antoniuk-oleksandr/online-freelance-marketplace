import type {User} from "@/types/User.ts";
import {request} from "@/api/request";
import {errorStore} from "@/common-stores/error-store.ts";
import { GetUserByIdRequestResponse } from "@/types/GetUserByIdRequestResponse";

export const tryToGetUserById = async (
    id: string,
    setUser: (newUser: User | null | undefined) => void
) => {

    const {data, status} = await request<GetUserByIdRequestResponse>(`/users/${id}`, 'GET')
    if (status !== 200) {
        errorStore.set({shown: true, error: data.error});
        return;
    }

    setUser(data);
}