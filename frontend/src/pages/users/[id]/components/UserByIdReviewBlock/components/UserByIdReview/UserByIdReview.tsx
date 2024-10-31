import UserByIdReviewLayout from "./UserByIdReviewLayout";
import {Review} from "@/types/Review";
import DividerElement from "@/common-components/DividerElement";
import UserByIdReviewUser
    from "@/pages/users/[id]/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewUser/UserByIdReviewUser";
import UserByIdReviewRatingBar
    from "@/pages/users/[id]/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewRatingBar/UserByIdReviewRatingBar";
import UserByIdReviewBottomContent
    from "@/pages/users/[id]/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewBottomContent/UserByIdReviewBottomContent";

type UserByIdReviewProps = {
    review: Review,
    showServices?: boolean
}

const UserByIdReview = (props: UserByIdReviewProps) => {
    const {review} = props;
    const {customer} = review;

    return (
        <UserByIdReviewLayout>
            <UserByIdReviewUser {...customer}/>
            <DividerElement/>
            <UserByIdReviewRatingBar {...review}/>
            <UserByIdReviewBottomContent {...props}/>
        </UserByIdReviewLayout>
    )
}

export default UserByIdReview;