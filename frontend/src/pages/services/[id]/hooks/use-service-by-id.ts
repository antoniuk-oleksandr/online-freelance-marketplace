import {useRouter} from "next/router";
import {useEffect, useState} from "react";
import {Service} from "@/types/Service";
import {getServiceByIdRequest} from "@/api/get-service-by-id-request";

export const useServiceById = () => {
    const [service, setService] = useState<Service | null>(null);
    const [status, setStatus] = useState<number | null>(null);
    const router = useRouter();

    useEffect(() => {
        if (!router.isReady) return;

        const getData = async () => {
            const id = router.query.id as string;
            const {data, status} = await getServiceByIdRequest(id);
            if (status === 200) {
                setService(data);
                setStatus(status);
            } else setStatus(status)
        }

        getData();
    }, [router]);

    return {service, status};
}