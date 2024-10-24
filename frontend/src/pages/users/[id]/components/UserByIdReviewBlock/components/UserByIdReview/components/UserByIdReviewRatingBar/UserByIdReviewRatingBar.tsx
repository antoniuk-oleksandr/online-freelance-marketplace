import UserByIdReviewRatingBarLayout from "./UserByIdReviewRatingBarLayout";
import {Rating} from "@mantine/core";
import BulletElement from "@/common-components/BulletElement";
import {getTimeAgo} from "@/utils/utils";
import {Review} from "@/types/Review";

type UserByIdReviewRatingBarProps = Review;

const UserByIdReviewRatingBar = (props: UserByIdReviewRatingBarProps) => {
    const {rating, endedAt} = props;

    return (
        <UserByIdReviewRatingBarLayout>
            <Rating color={"rgb(250, 204, 21)"} value={rating} readOnly/>
            <p className={"text-base font-semibold"}>{rating}</p>
            <BulletElement/>
            <p
                className={"text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}
            >{getTimeAgo(endedAt)}</p>
        </UserByIdReviewRatingBarLayout>
    )
}

export default UserByIdReviewRatingBar;