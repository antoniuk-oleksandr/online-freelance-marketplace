import {LayoutProps} from "@/types/LayoutProps";

type UserInfoReviewsLayoutProps = LayoutProps & {
    size?: "small" | "large",
};

const UserInfoReviewsLayout = (props: UserInfoReviewsLayoutProps) => {
    const {children, size} = props;

    return (
        <div
            className={`flex items-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary gap-x-2 
            ${size === "small" ? "text-base" : "text-lg"}`}
        >
            {children}
        </div>
    )
}

export default UserInfoReviewsLayout;