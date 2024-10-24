import {useEffect, useState} from "react";
import {useRouter} from "next/router";
import {getUserByIdRequest} from "@/api/get-user-by-id-request";
import {User} from "@/types/User";

export const useUserById = () => {
    const [user, setUser] = useState<User | null>(null);
    const [status, setStatus] = useState<null | number>(null);
    const router = useRouter();

    useEffect(() => {
        if(!router.isReady) return;

        const {id} = router.query;

        const getData = async () => {
            const {data, status} = await getUserByIdRequest(id as string);
            if(status === 200) setUser(data);
            else setStatus(status);
        }

        getData();
    }, [router])

    return {user, status};
}