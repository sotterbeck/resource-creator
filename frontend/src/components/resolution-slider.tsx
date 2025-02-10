import { Slider } from "@/components/ui/slider";

interface ResolutionSliderProps {
  max: number;
  min: number;
  value: number;
  onValueChange: (value: number) => void;
}

export function ResolutionSlider({
  onValueChange,
  value,
  max,
  min,
}: ResolutionSliderProps) {
  const maxExponent = Math.floor(Math.log2(max));
  const minExponent = Math.floor(Math.log2(min));
  const currentExponent = Math.log2(value);

  return (
    <div className="space-y-2">
      <div className={"flex justify-between"}>
        <span className="text-sm">Resolution</span>
        <span className="text-nowrap font-mono text-sm tabular-nums text-muted-foreground">{`${value} px`}</span>
      </div>
      <Slider
        value={[currentExponent]}
        onValueChange={([exponent]) => onValueChange(2 ** exponent)}
        step={1}
        min={minExponent}
        max={maxExponent}
        disabled={!max}
      />
    </div>
  );
}
