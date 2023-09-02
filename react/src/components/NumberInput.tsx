

type PropsType = {
    name: string
    nameText: string
    BoolValue: boolean
    onChange: (checked: boolean) => void;
}

const NumberInput = (props: PropsType) => {
    return (
        <div className="flex flex-row w-2/3 m-1 p-2 pl-8 font-semibold rounded-md shadow-md bg-yellow-300 justify-start items-center">
            <label htmlFor="Port" className="w-32 border border-black rounded-md">
                {props.nameText}:
            </label>
            <input
                className="rounded ml-2 px-1"
                {...register("Port", { min: 0, max: 65535 })}
                id="Port"
                type="number"
                value={config.Port}
                onChange={onChange}
                placeholder="Port"
            />
            <div className="bg-red-600">
                {errors.Port && <div>入力が必須の項目です(0~65535)</div>}
            </div>
        </div>
    );
};

export default NumberInput;