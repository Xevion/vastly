import { api } from "@wails/go/models";

export default function Offer({
  offer: scoredOffer,
}: {
  offer: api.ScoredOffer;
}) {
  return <div className="p-4">{scoredOffer.Score}</div>;
}
