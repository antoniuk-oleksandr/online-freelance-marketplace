import UserInfoReviewsLayout from "./UserInfoReviewsLayout";
import {FaStar} from "react-icons/fa6";
import {UserByIdProps} from "@/types/UserByIdProps";

type UserInfoReviewsProps = UserByIdProps & {
    size?: "small" | "large",
};

const UserInfoReviews = (props: UserInfoReviewsProps) => {
    const {user} = props;
    const {rating, reviewsCount} = user;

    return (
        <UserInfoReviewsLayout {...props}>
            <FaStar className="text-yellow-400"/>
            <span className="font-semibold">{rating}</span>
            <span>({reviewsCount})</span>
        </UserInfoReviewsLayout>
    )
}

export default UserInfoReviews;