import {request} from "@/api/request.ts";
import {errorStore} from "@/common-stores/error-store.ts";
import type {Package} from "@/types/Package.ts";
import type {GetRestrictedServiceByIdResponse} from "@/types/GetRestrictedServiceByIdResponse.ts";
import {RestrictedService} from "@/types/RestrictedService.ts";
import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";

export const parseOrderRequestQuery = (
    setServiceId: (value: string) => void,
    setPackageId: (value: string) => void
) => {
    const packageId = new URLSearchParams(window.location.search).get("packageId");
    const serviceId = new URLSearchParams(window.location.search).get("serviceId");
    if (!packageId || !serviceId) {
        errorStore.set({shown: true, error: ResponseErrorEnum.UnexpectedError})
        return;
    }

    setServiceId(serviceId);
    setPackageId(packageId);
}

export const fetchServiceDetailsAndPackage = (
    serviceId: string,
    packageId: string,
    setServiceData: (data: RestrictedService) => void,
    setSelectedPackage: (newPackage: Package) => void
) => {
    request<GetRestrictedServiceByIdResponse>("GET", `/freelances/${serviceId}/restricted/`).then((response) => {
        if (response.status === 200) {
            const data = response.data;
            setServiceData(data);

            if (!packageId) setSelectedPackage(response.data.packages[0]);
            else {
                const selectedPackage = data.packages.find((pkg) => pkg.id === parseInt(packageId))
                    || data.packages[0];
                setSelectedPackage(selectedPackage);
            }
        } else errorStore.set({shown: true, error: response.data.error});
    });
}
