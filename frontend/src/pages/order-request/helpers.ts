import {request} from "@/api/request.ts";
import type {GetServiceByUserIdRequestResponse} from "@/types/GetServiceByIdRequestResponse.ts";
import {errorStore} from "@/common-stores/error-store.ts";
import {Service} from "@/types/Service.ts";
import type {Package} from "@/types/Package.ts";

export const fetchServiceDetailsAndPackage = (
    serviceId: string,
    setServiceData: (data: Service) => void,
    setSelectedPackage: (newPackage: Package) => void
) => {
    const packageId = new URLSearchParams(window.location.search).get("packageId");

    request<GetServiceByUserIdRequestResponse>(`/freelances/${serviceId}`, "GET").then((response) => {
        if (response.status === 200) {
            const data = response.data.service;
            setServiceData(data);

            if (!packageId) setSelectedPackage(response.data.service.packages[0]);
            else {
                const selectedPackage = data.packages.find((pkg) => pkg.id === parseInt(packageId))
                    || data.packages[0];
                setSelectedPackage(selectedPackage);
            }
        } else errorStore.set({shown: true, error: response.data.error});
    });
}