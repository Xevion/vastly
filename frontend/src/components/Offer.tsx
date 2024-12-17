import { api } from "@wails/go/models";
import { useState } from "react";
import { cn } from "@src/utils";
import { Tooltip } from "react-tooltip";

export default function Offer({
  offer: { Offer: offer, Score: score, Reasons: reasons },
}: {
  offer: api.ScoredOffer;
}) {
  const copy = (text: string) => navigator.clipboard.writeText(text);
  const [showDetails, setShowDetails] = useState(false);
  const mb_to_gb = (mb: number) => Math.round(mb / 1024);

  return (
    <div
      className={cn({
        "flex [&>*]:px-2 flex-col relative bg-zinc-700/90 rounded max-w-md overflow-hidden":
          true,
        "h-24": !showDetails,
        "min-h-24 max-h-48": showDetails,
      })}
    >
      <div className="flex">
        <span className="text-4xl font-bold pr-2">
          {score >= 10 ? Math.round(score) : score.toFixed(1)}
        </span>
        <span className="relative text-xl top-2.5">
          {offer.num_gpus}x {offer.gpu_name}{" "}
          <span className="text-sm">{mb_to_gb(offer.gpu_ram)} GB</span>
        </span>
      </div>
      <div>
        <span>${offer.search.totalHour.toFixed(2)}/hr</span>
        <span className="pl-3 text-sm">
          <span className="text-xs select-none">mem</span>{" "}
          {mb_to_gb(offer.cpu_ram)}/{mb_to_gb(offer.cpu_ram / offer.gpu_frac)}GB
        </span>
        <span className="pl-3 text-sm">
          <span className="text-xs select-none">dlperf</span>{" "}
          {offer.dlperf.toFixed(0)}
        </span>
      </div>
      <div className="select-none [>button]:select-auto [>button]:text-blue-500 w-full left-1 text-xs space-x-1">
        <button onClick={() => copy(offer.machine_id.toString())}>
          m{offer.machine_id}
        </button>
        <button onClick={() => copy(offer.host_id.toString())}>
          h{offer.host_id}
        </button>
        <span>{Math.round(offer.duration / 60 / 60 / 24)} days</span>
        <span
          className={
            offer.verification != "verified" ? "text-orange-400/90" : ""
          }
        >
          {offer.verification}
        </span>
      </div>
      <div
        onClick={() => setShowDetails(!showDetails)}
        className={cn({
          "px-0 w-full bg-zinc-900/70 border-t cursor-pointer border-zinc-600/80   text-center":
            true,
          "select-none h-3 leading-[0.2rem] text-zinc-100 absolute bottom-0":
            !showDetails,
          "h-40 overflow-y-auto text-sm": showDetails,
        })}
      >
        {showDetails ? (
          <>
            <Tooltip id="reason" />
            {reasons
              .sort((a, _) => (a.IsMultiplier ? 1 : -1))
              .map((reason, i) => (
                <div
                  data-tooltip-id="reason"
                  data-tooltip-content={reason.Value}
                  key={i}
                  className={cn(
                    "space-x-2",
                    (reason.IsMultiplier && reason.Offset < 1) ||
                      (!reason.IsMultiplier && reason.Offset < 0)
                      ? ""
                      : ""
                  )}
                >
                  {reason.IsMultiplier ? (
                    <span>x{reason.Offset.toFixed(2)}</span>
                  ) : (
                    <span>
                      {reason.Offset > 0 ? "+" : null}
                      {reason.Offset}
                    </span>
                  )}
                  <span>{reason.Reason}</span>
                </div>
              ))}
          </>
        ) : (
          "..."
        )}
      </div>
    </div>
  );
}
