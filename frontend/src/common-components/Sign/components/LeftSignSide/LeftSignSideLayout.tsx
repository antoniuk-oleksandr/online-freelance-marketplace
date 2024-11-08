import {LayoutProps} from "@/types/LayoutProps";

const LeftSignSideLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"flex flex-col justify-center gap-y-8 px-4 md:px-24"}>
            {children}
        </div>
    )
}

export default LeftSignSideLayout;