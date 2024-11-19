import UserByIdReviewServiceLayout from "./UserByIdReviewServiceLayout";
import {UserReview} from "@/types/UserReview";
import UserByIdReviewServiceImage
    from "@/pages/users/[id]/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewServiceImage/UserByIdReviewServiceImage";

type UserByIdReviewServiceProps = UserReview;

const UserByIdReviewService = (props: UserByIdReviewServiceProps) => {
    const {image, title} = props;

    return (
        <UserByIdReviewServiceLayout {...props}>
            <UserByIdReviewServiceImage src={image}/>
            <span className={"line-clamp-2 w-auto md:w-40"}>{title}</span>
        </UserByIdReviewServiceLayout>
    )
}

export default UserByIdReviewService;