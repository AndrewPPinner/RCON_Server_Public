export class Players {
	Name:    string
	UID:     string
	SteamID: string

    constructor(name: string, uid: string, steamId: string){
        this.Name = name;
        this.UID = uid;
        this.SteamID = steamId;
    }
}