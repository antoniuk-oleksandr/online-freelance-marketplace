import {Review} from "@/types/Review";

type UserByIdReviewContentProps = Review;

const UserByIdReviewContent = (props: UserByIdReviewContentProps) => {
    const {content} = props;

    return (
        <div className={"text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}>
            {content}
        </div>
    )
}

export default UserByIdReviewContent;