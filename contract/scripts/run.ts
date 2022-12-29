import { ethers } from "hardhat";

async function main() {

  const contractFactory = await ethers.getContractFactory("Elections");
  const electionsContract = await contractFactory.deploy();

  await electionsContract.deployed();

  console.log("Elections contract deployed to:", electionsContract.address);

  await electionsContract.createElection("Test Election","some-random-uuid");
  console.log("Election created");

  const election = await electionsContract.getElectionById("some-random-uuid");
  console.log("Election retrieved", election);

  await electionsContract.addCandidate("nico","1","some-random-uuid");
  await electionsContract.addCandidate("nico2","2","some-random-uuid");

  await electionsContract.startElection("some-random-uuid");
  console.log("Election started");
  
  await electionsContract.vote("1","some-random-uuid");
  await electionsContract.vote("1","some-random-uuid");
  await electionsContract.vote("1","some-random-uuid");
  await electionsContract.vote("2","some-random-uuid");
  await electionsContract.vote("2","some-random-uuid");

  const candidates = await electionsContract.getAllCandidates();
  console.log("Candidates retrieved", candidates);

  await electionsContract.finishElection("some-random-uuid");

  const winner = await electionsContract.getWinner("some-random-uuid");
  console.log("Winner retrieved", winner);
}


// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });