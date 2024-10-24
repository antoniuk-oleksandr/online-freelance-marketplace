type DividerElementProps = {
    orientation?: "horizontal" | "vertical",
    className?: string,
};

const DividerElement = (props: DividerElementProps) => {
    const {orientation, className} = props;

    let styles = '';
    if(orientation ==="horizontal" || !orientation) styles = 'border-t w-full';
    else styles = 'border-l h-full';

    return <div className={`${className} ${styles} border-light-palette-divider dark:border-dark-palette-divider`}></div>
}

export default DividerElement;