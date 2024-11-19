import UserByIdReviewColumnLayout from "./UserByIdReviewColumnLayout";

type UserByIdReviewColumnProps = {
    label: string,
    data: string | number
}

const UserByIdReviewColumn = (props: UserByIdReviewColumnProps) => {
    const {label, data} = props;

    return (
        <UserByIdReviewColumnLayout>
            <p className={"font-semibold text-base"}>{data}</p>
            <p className={"text-light-palette-text-secondary dark:text-dark-palette-text-secondary w-20 md:w-fit"}>{label}</p>
        </UserByIdReviewColumnLayout>
    )
}

export default UserByIdReviewColumn;