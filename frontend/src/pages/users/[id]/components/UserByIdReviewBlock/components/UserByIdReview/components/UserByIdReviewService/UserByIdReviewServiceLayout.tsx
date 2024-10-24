import {LayoutProps} from "@/types/LayoutProps";
import Link from "next/link";
import {UserReview} from "@/types/UserReview";

type UserByIdReviewServiceLayoutPros = LayoutProps & UserReview;

const UserByIdReviewServiceLayout = (props: UserByIdReviewServiceLayoutPros) => {
    const {children, id} = props;

    return (
        <Link
            href={`/services/${id}`}
            className={"py-2 px-4 ring-1 ring-light-palette-divider shadow rounded-xl flex gap-4 items-center w-full md:w-fit"}
        >
            {children}
        </Link>
    )
}

export default UserByIdReviewServiceLayout;