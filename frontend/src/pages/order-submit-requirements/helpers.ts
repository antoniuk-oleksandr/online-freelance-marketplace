import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import type { GetPublicKeyRequestResponse } from "@/types/GetOrderServiceQuestionsRequestResponse";
import type { OrderSubmitRequirementsData } from "@/types/OrderSubmitRequirementsData";
import type { OrderSubmitRequirementsFormErrors } from "@/types/OrderSubmitRequirementsFormErrors";
import type { ServiceQuestion } from "@/types/ServiceQuestion";
import { StatusEnum } from "@/types/StatusEnum";
import Cookies from "js-cookie";
import { z } from "zod";

export const getSubmitRequirementsSteps = (
    orderData: OrderSubmitRequirementsData | undefined
) => {
    if (!orderData) return undefined;

    return [
        {
            text: "Order details",
            link: `/orders/request?serviceId=${orderData.service.id}&packageId=${orderData.service.package.id}`,
        },
        {
            text: "Confirm & pay",
            link: `/orders/confirm-pay?serviceId=${orderData.service.id}&packageId=${orderData.service.package.id}`
        },
        { text: "Submit requirements" }
    ];
}

export const makeOrderServiceQuestionsRequest = (
    orderId: string,
    setOrderData: (value: OrderSubmitRequirementsData) => void
) => {
    const token = Cookies.get("accessToken");

    request<GetPublicKeyRequestResponse>(`/orders/${orderId}/freelance-questions`, "GET", token)
        .then((response) => {
            if (response.status !== 200) {
                errorStore.set({ shown: true, error: response.data.error })
            } else setOrderData(response.data);

        });
}

export const formatOrderDate = (dateStr: string): string => {
    const date = new Date(dateStr);

    const month = getMonthAbbreviation(date.getUTCMonth())
    const day = date.getDate();
    const year = date.getUTCFullYear();

    return `${month} ${day},  ${year}`;
}

const getMonthAbbreviation = (monthIndex: number) => {
    return new Date(2000, monthIndex, 1).toLocaleString('en', { month: 'short' });
};

export const getStatusBgColor = (num: StatusEnum) => {
    const statusBgColorArr: Record<StatusEnum, string> = {
        [StatusEnum.Incomplete]: "bg-orange-500",
        [StatusEnum.AwaitingAcceptance]: "bg-orange-500",
        [StatusEnum.InProgress]: "bg-orange-500",
        [StatusEnum.Completed]: "bg-green-500",
        [StatusEnum.Cancelled]: "bg-red-500",
        [StatusEnum.Pending]: "bg-orange-500",
        [StatusEnum.Failed]: "bg-red-500",
    };

    return statusBgColorArr[num] || "text-gray-500 bg-gray-100";
}

export const orderSubmitRequirementsFormSchema = z.object({
    customerMessage: z.string()
        .min(1, { message: "Message is required." })
        .max(2000, { message: "Message is too long." }),
    answers: z.array(z.object({
        questionId: z.number(),
        content: z.string()
            .min(1, { message: "Answer is required." })
            .max(2000, { message: "Answer is too long." })
    }))
});

export const formatOrderFormData = (
    data: any
): OrderSubmitRequirementsFormErrors | undefined => {
    let customerMessage = "";
    let answers: string[] = [];

    if (data && data.customerMessage) {
        customerMessage = data.customerMessage[0];
    }
    if (data && data.answers) {
        answers = data.answers.map((answer: any) => {
            return answer.content ? answer.content[0] : undefined;
        });
    }

    return {
        customerMessage,
        answers,
    }
}

export const getOrderSubmitRequirementsFormInitialValues = (
    freelanceQuestions: ServiceQuestion[],
    orderId: string
) => ({
    answers: freelanceQuestions.map((question) => ({
        questionId: question.id,
        content: '',
    })),
    customerMessage: '',
    files: [],
    orderId: orderId,
})

