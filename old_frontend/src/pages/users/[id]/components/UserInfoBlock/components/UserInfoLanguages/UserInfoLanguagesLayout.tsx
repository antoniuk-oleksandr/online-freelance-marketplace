import {LayoutProps} from "@/types/LayoutProps";

const UserInfoLanguagesLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div
            className="flex items-center text-lg text-light-palette-text-secondary dark:text-dark-palette-text-secondary gap-x-2"
        >
            {children}
        </div>
    )
}

export default UserInfoLanguagesLayout;