import {ServiceQuestion} from "@/types/ServiceQuestion";
import {StatusEnum} from "@/types/StatusEnum";

export type OrderSubmitRequirementsData = {
    freelanceQuestions: ServiceQuestion[],
    order: {
        id: number,
        createdAt: string,
        status: StatusEnum,
    },
    service: {
        id: number,
        image: string,
        package: {
            id: number,
            price: number,
            title: string,
        },
        title: string,
    },
    files: File[],
}
