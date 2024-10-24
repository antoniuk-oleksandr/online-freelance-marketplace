import {LayoutProps} from "@/types/LayoutProps";

type PaperElementProps = LayoutProps & {
    styles?: string,
}

const PaperElement = (props: PaperElementProps) => {
    const {children, styles} = props;

    return (
        <div
            className={`${styles} py-4 px-6 rounded-lg border border-light-palette-divider dark:border-dark-palette-divider`}
        >
            {children}
        </div>
    )
}

export default PaperElement;