class Pair {
    constructor(key, value) {
        this.key = key;
        this.value = value;
        this.next = null;
    }
}

class HashTable {
    constructor(size) {
        this.cells = new Array(size);
    }

    hash(key) {
        if (key === null || key === undefined) {
            return 0; // Could be any other default hash value, using 0 for simplicity
        }
    
        let total = 0;
        for (let i = 0; i < key.length; i++) {
            total += key.charCodeAt(i);
        }
    
        return total % this.cells.length;
    } 

    insert(key, value) {
        const hash = this.hash(key);

        if (!this.cells[hash]) {
            this.cells[hash] = new Pair(key, value);
        } else if (this.cells[hash].key === key) {
            this.cells[hash].value = value;
        } else {
            let keyValuePair = this.cells[hash];

            while (keyValuePair.next) {
                if (keyValuePair.next.key === key) {
                    keyValuePair.next.value = value;
                    return;
                }

                keyValuePair = keyValuePair.next;
            }

            keyValuePair.next = new Pair(key, value);
        }
    }

    get(key) {
        const hash = this.hash(key);
        if (!this.cells[hash]) {
            return null;
        }

        let keyValuePair = this.cells[hash];

        while (keyValuePair) {
            if (keyValuePair.key === key) {
                return keyValuePair.value;
            }

            keyValuePair = keyValuePair.next;
        }
    }

    getAll() {
        const table = [];
    
        for (let i = 0; i < this.cells.length; i++) {
            const cell = [];
            let keyValuePair = this.cells[i];
    
            while (keyValuePair) {
                cell.push([keyValuePair.key, keyValuePair.value]);
                keyValuePair = keyValuePair.next;
            }
    
            table.push(cell);
        }
    
        return table;
    }    
    
}

module.exports = HashTable;