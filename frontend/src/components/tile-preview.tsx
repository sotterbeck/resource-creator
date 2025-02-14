interface TilePreviewProps {
  src: string;
  patternResolution: { width: number; height: number };
  tileResolution: number;
}

export function TilePreview({
  src,
  tileResolution,
  patternResolution,
}: TilePreviewProps) {
  if (!src) {
    return <></>;
  }

  const cellsX = Math.floor(patternResolution.width / tileResolution);
  const cellsY = Math.floor(patternResolution.height / tileResolution);
  const cellWidth = `${100 / cellsX}%`;
  const cellHeight = `${100 / cellsY}%`;

  return (
    <div className="relative w-full overflow-hidden rounded-md">
      <img
        src={src}
        alt="A preview of the selected texture"
        className="pixelated h-full w-full object-cover"
      />
      <div
        className="pointer-events-none absolute inset-0"
        style={{
          backgroundImage: `
            linear-gradient(to right, rgba(255,255,255,0.7) 1px, transparent 1px),
            linear-gradient(to bottom, rgba(255,255,255,0.7) 1px, transparent 1px)
          `,
          backgroundSize: `${cellWidth} ${cellHeight}`,
        }}
      />
    </div>
  );
}
