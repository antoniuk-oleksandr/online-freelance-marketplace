import {LayoutProps} from "@/types/LayoutProps";

const UserByIdReviewColumnLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"md:w-20  flex items-center flex-row-reverse md:block"}>
            {children}
        </div>
    )
}

export default UserByIdReviewColumnLayout;