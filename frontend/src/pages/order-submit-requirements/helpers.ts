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
    const statusClasses = {
        [StatusEnum.Incomplete]: "bg-gray-200 text-gray-700 border-gray-300",
        [StatusEnum.InProgress]: "bg-yellow-100 text-yellow-700 border-yellow-300",
        [StatusEnum.Completed]: "bg-green-100 text-green-700 border-green-300",
        [StatusEnum.Cancelled]: "bg-red-200 text-red-800 border-red-400",
        [StatusEnum.Pending]: "bg-orange-100 text-orange-700 border-orange-300",
        [StatusEnum.Failed]: "bg-red-100 text-red-700 border-red-300",
        [StatusEnum.AwaitingAcceptance]: "bg-blue-100 text-blue-700 border-blue-300"
    };

    return statusClasses[num] || "text-gray-500 bg-gray-100";

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

