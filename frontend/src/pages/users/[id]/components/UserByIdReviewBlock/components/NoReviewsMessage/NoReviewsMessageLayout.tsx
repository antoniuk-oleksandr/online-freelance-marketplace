import {LayoutProps} from "@/types/LayoutProps";

const NoReviewsMessageLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"text-center text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}>
            {children}
        </div>
    )
}

export default NoReviewsMessageLayout;