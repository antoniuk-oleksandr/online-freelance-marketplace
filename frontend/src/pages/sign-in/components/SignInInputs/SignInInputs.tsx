import MyInput from "@/common-components/MyInput/MyInput";
import {useFormContext} from "react-hook-form";
import {getErrorMessage} from "@/utils/input-utils";

const SignInInputs = () => {
    const {formState, register} = useFormContext();
    const {errors} = formState;

    return (
        <>
            <MyInput
                register={register}
                name={"email"}
                error={getErrorMessage(errors, 'email')}
                label={"Email"}/>
            <MyInput
                inputType={"password"}
                register={register}
                name={"password"}
                error={getErrorMessage(errors, 'password')}
                label={"Password"}
            />
        </>
    )
}

export default SignInInputs;