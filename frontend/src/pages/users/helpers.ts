import {getUserByIdRequest} from "@/api/get-user-by-id-request.ts";
import {getFile} from "@/utils/utils.ts";
import type {UserService} from "@/types/UserService.ts";
import type {User} from "@/types/User.ts";

export const tryToGetUserById = async (
    id: string,
    setUser: (newUser: User | null | undefined) => void
) => {
    const {data, status} = await getUserByIdRequest(id);
    if (status === 404) setUser(null);
    if (!data) return;

    data.avatar = getFile(data.avatar);
    data.services && data.services.map((service: UserService) => {
        service.image = getFile(service.image);
    });
    data.reviews && data.reviews.map((review) => {
        review.service.image = getFile(review.service.image);
        review.customer.avatar = getFile(review.customer.avatar);
    });

    setUser(data);
}