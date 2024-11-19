import {LayoutProps} from "@/types/LayoutProps";

const UserComponentLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"flex flex-col gap-y-2"}>
            {children}
        </div>
    )
}

export default UserComponentLayout;