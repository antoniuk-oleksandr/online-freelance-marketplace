import type {OrderParams} from "@/types/OrderParams.ts";

export const getPaymentPageParams = (): null | OrderParams => {
    const pageParams = new URLSearchParams(window.location.search);
    const paramNames = ["serviceID", "packageID", "message"] as unknown as (keyof OrderParams)[];

    return paramNames.reduce((acc, name) => {
        const value = pageParams.get(name);

        if (!value) acc = null;
        else if (value && acc) acc[name] = value;

        return acc;
    }, {} as OrderParams | null);
}