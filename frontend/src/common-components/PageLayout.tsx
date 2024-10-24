import {LayoutProps} from "@/types/LayoutProps";

const PageLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div
            className={"mx-auto px-4 w-full md:max-w-256 xl:max-w-320"}>
            {children}
        </div>
    )
}

export default PageLayout;