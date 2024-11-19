import {Loader} from "@mantine/core";
import {useEffect, useState} from "react";

const AppLoader = () => {
    const [ready, setReady] = useState(false);

    useEffect(() => {
        const timeout = setTimeout(() => {
            setReady(true);
        }, 5000);

        return () => clearTimeout(timeout);
    }, []);

    if(!ready) return null;
    return (
        <div
            className={"absolute grid place-items-center left-0 top-0 bg-gray-100 bg-opacity-90 w-svw h-svh dark:bg-black dark:bg-opacity-90"}>
            <Loader color={"blue"} type={"dots"} size={50}/>
        </div>
    )
}

export default AppLoader;