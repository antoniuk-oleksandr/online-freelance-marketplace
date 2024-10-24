import {LayoutProps} from "@/types/LayoutProps";

const ServiceAboutBlockLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"flex flex-col gap-y-2 text-base"}>
            {children}
        </div>
    )
}

export default ServiceAboutBlockLayout;