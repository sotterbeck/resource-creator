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

  const tilesX = Math.floor(patternResolution.width / tileResolution);
  const tilesY = Math.floor(patternResolution.height / tileResolution);

  return (
    <div className="relative w-full overflow-hidden rounded-md">
      <img
        src={src}
        alt="A preview of the selected texture"
        className="h-full w-full object-cover pixelated"
      />
      <div
        className="absolute inset-0 grid w-full gap-0"
        style={{
          gridTemplateColumns: `repeat(${tilesX}, 1fr)`,
          gridTemplateRows: `repeat(${tilesY}, 1fr)`,
        }}
      >
        {Array.from({ length: tilesX * tilesY }).map((_, i) => (
          <div key={i} className="border border-gray-400"></div>
        ))}
      </div>
    </div>
  );
}
