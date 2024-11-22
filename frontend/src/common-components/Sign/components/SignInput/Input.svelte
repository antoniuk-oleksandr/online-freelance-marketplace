<script lang="ts">
    import InputLayout from "@/common-components/Sign/components/SignInput/InputLayout.svelte";
    import InputError from "@/common-components/Sign/components/SignInput/components/InputError/InputError.svelte";
    import InputLabel from "@/common-components/Sign/components/SignInput/components/InputLabel/InputLabel.svelte";
    import InputElement from "@/common-components/Sign/components/SignInput/components/InputElement/InputElement.svelte";
    import InputContainerLayout from "@/common-components/Sign/components/SignInput/components/InputContainer/InputContainerLayout.svelte";
    import InputPasswordIcon from "@/common-components/Sign/components/SignInput/components/InputPasswordIcon/InputPasswordIcon.svelte";

    type InputProps = {
        id: string,
        type: "text" | "password",
        label: string,
        value: string,
        error: string | null,
        wasSubmitted: boolean,
    }

    let {id, type, label, value, error, wasSubmitted}: InputProps = $props();

    let isFocused = $state(false);
    const setIsFocused = (value: boolean) => isFocused = value;

    let showPassword = $state(false);
    const setShowPassword = (value: boolean) => showPassword = value;
</script>

<InputLayout>
    <InputContainerLayout
            isFocused={isFocused} error={error}
            wasSubmitted={wasSubmitted} value={value}
    >
        <InputLabel
                id={id} value={value}
                label={label} error={error}
                wasSubmitted={wasSubmitted}
                isFocused={isFocused}
        />
        <InputElement
                id={id} type={type}
                setIsFocused={setIsFocused}
                showPassword={showPassword}
        />
        <InputPasswordIcon
                showPassword={showPassword}
                setShowPassword={setShowPassword}
                type={type}/>
    </InputContainerLayout>
    <InputError wasSubmitted={wasSubmitted} error={error}/>
</InputLayout>

