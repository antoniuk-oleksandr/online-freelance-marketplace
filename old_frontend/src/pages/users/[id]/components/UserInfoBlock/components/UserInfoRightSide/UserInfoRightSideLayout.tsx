import {LayoutProps} from "@/types/LayoutProps";

const UserInfoRightSideLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className="flex flex-col justify-center">
            {children}
        </div>
    )
}

export default UserInfoRightSideLayout;