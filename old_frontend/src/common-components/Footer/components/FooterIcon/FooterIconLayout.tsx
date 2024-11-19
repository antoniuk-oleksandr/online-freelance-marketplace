import {LayoutProps} from "@/types/LayoutProps";

type FooterIconProps = LayoutProps & {
    link: string,
}

const FooterIconLayout = (props: FooterIconProps) => {
    const {children, link} = props;

    return (
        <div
            onClick={() => window.open(props.link, "_blank")}
            className={"cursor-pointer ring-light-palette-action-hover dark:ring-dark-palette-action-hover hover:ring-8 duration-200 ease-out hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover rounded-full"}
        >
            {children}
        </div>
    )
}

export default FooterIconLayout;