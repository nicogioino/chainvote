//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

contract Elections {
    // Struct to represent a candidate
    struct Candidate {
        string name;
        string id;
        uint voteCount;
    }

    // Mapping from candidate IDs to candidates
    mapping(string => Candidate) public candidates;

    // Array of all candidate IDs
    string[] public candidateIds;

    // Enum to represent the status of an election
    enum ElectionStatus {
        NOT_STARTED,
        IN_PROGRESS,
        FINISHED
    }

    // Struct to represent an election
    struct Election {
        string name;
        ElectionStatus status;
        string[] candidateIds;
        string id;
        address owner;
    }

    // Mapping from election ids to elections
    mapping(string => Election) public elections;

    // Create a new election with an empty list of candidates
    function createElection(string memory name, string memory id) public {
        // Ensure that the election does not already exist
        require(elections[id].status == ElectionStatus.NOT_STARTED, "Election already exists");

        // Create the new election
        elections[id] = Election({
            name: name,
            status: ElectionStatus.NOT_STARTED,
            id: id,
            candidateIds: new string[](0),
            owner: msg.sender
        });
    }

    // Add a candidate to an election
    function addCandidate(string memory name,string memory id,string memory electionId) public {
    
        require(msg.sender == elections[electionId].owner, "Only the owner can add candidates");

        // Ensure that the candidate does not already exist
        require(candidates[id].voteCount == 0, "Candidate already exists");

        //Ensure that the election hasn't started
        require(
            elections[electionId].status == ElectionStatus.NOT_STARTED,
            "Election has already started"
        );

        // Add the candidate to the list of candidates
        candidates[id] = Candidate({name: name, id: id, voteCount: 0});

        // Add the candidate ID to the list of candidate IDs
        candidateIds.push(id);

        // Add the candidate ID to the election
        elections[electionId].candidateIds.push(id);
    }

    // Cast a vote in a given election
    function vote(string memory candidateId, string memory electionId) public {
        // Ensure that the election is in progress
        Election storage election = elections[electionId];
        require(
            election.status == ElectionStatus.IN_PROGRESS,
            "Election is not in progress"
        );

        // Increment the vote count for the candidate
        candidates[candidateId].voteCount++;
    }

    // Finish an election
    function finishElection(string memory id) public {
        require(msg.sender == elections[id].owner, "Only the owner can finish the election");
        // Ensure that the election is in progress
        Election storage election = elections[id];
        require(
            election.status == ElectionStatus.IN_PROGRESS,
            "Election is not in progress"
        );

        // Set the election status to finished
        election.status = ElectionStatus.FINISHED;

        //save the new status
        elections[id] = election;
    }

    // Get the winner of an election
    function getWinner(string memory name) public view returns (string memory) {
        // Ensure that the election is finished
        Election storage election = elections[name];
        require(election.status == ElectionStatus.FINISHED,"Election is not finished");

        // Initialize the winner to the first candidate
        string memory winner = election.candidateIds[0];

        // Iterate over the candidates and find the one with the most votes
        for (uint i = 1; i < election.candidateIds.length; i++) {
            string memory candidateId = election.candidateIds[i];
            if (
                candidates[candidateId].voteCount > candidates[winner].voteCount
            ) {
                winner = candidateId;
            }
        }

        return winner;
    }

    function getElectionById(string memory id) public view returns (Election memory) {
        return elections[id];
    }

    function getAllCandidates() public view returns (Candidate[] memory) {
        Candidate[] memory allCandidates = new Candidate[](candidateIds.length);
        for (uint i = 0; i < candidateIds.length; i++) {
            allCandidates[i] = candidates[candidateIds[i]];
        }
        return allCandidates;
    }

    function compareStrings(string memory a,string memory b) public pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
    }

    function startElection(string memory id) public {
        require(msg.sender == elections[id].owner, "Only the owner can start the election");
        // Ensure that the election is not in progress
        Election storage election = elections[id];
        require(
            election.status == ElectionStatus.NOT_STARTED,
            "Election has already started"
        );

        election.status = ElectionStatus.IN_PROGRESS;
        elections[id] = election;
    }
}
