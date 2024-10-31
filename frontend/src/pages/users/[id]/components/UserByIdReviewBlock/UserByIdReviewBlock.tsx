import UserComponentLayout from "@/pages/users/[id]/components/UserComponentLayout";
import UserByIdReview from "@/pages/users/[id]/components/UserByIdReviewBlock/components/UserByIdReview/UserByIdReview";
import {Review} from "@/types/Review";
import NoReviewsMessage
    from "@/pages/users/[id]/components/UserByIdReviewBlock/components/NoReviewsMessage/NoReviewsMessage";

type UserByIdReviewProps = {
    reviews: Review[],
    showServices?: boolean
}

const UserByIdReviewBlock = (props: UserByIdReviewProps) => {
    const {reviews} = props;

    if (!reviews) return null;

    return (
        <UserComponentLayout>
            <p className={"text-xl font-bold"}>Reviews</p>
            <div className={"flex flex-col gap-y-4"}>
                {reviews.length === 0
                    ? <NoReviewsMessage/>
                    : reviews.map((review, index) => (
                        <UserByIdReview
                            {...props}
                            key={index}
                            review={review}
                        />
                    ))
                }
            </div>
        </UserComponentLayout>
    )
}

export default UserByIdReviewBlock;