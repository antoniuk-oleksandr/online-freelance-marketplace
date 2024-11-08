import {useEffect} from "react";
import {useRouter} from "next/router";

export const usePage = () => {
    const router = useRouter();

    useEffect(() => {
        if(!router.isReady) return;
    }, [router]);

    return {page: router.pathname}
}