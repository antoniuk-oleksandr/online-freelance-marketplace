import {LayoutProps} from "@/types/LayoutProps";

const UserByIdServiceInfoLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div
            className="text-base flex items-center mb-2 text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
        >
            {children}
        </div>
    )
}

export default UserByIdServiceInfoLayout;