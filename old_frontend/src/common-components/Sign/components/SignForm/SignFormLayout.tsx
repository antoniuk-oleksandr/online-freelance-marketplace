import {LayoutProps} from "@/types/LayoutProps";
import {FormProvider, useForm} from "react-hook-form";
import {SignProps} from "@/types/SignProps";
import {zodResolver} from "@hookform/resolvers/zod";

type SignFormLayoutProps = LayoutProps & SignProps;

const SignFormLayout = (props: SignFormLayoutProps) => {
    const {children, onSubmit, schema, defaultValues} = props;
    const methods = useForm({
        defaultValues,
        reValidateMode: "onChange",
        resolver: zodResolver(schema),
    });

    return (
        <FormProvider {...methods}>
            <form
                onSubmit={methods.handleSubmit(onSubmit)}
                className={"flex flex-col gap-y-5"}>
                {children}
            </form>
        </FormProvider>
    )
}

export default SignFormLayout;