class Trie{
public:
    bool isWord;
    Trie* children[26];
    
    Trie(){
        this->isWord = false;
        for(int i=0;i!=26;++i){
            this->children[i] = nullptr;
        }
    }
    
    void insert(string word){
        Trie* current = this;
        for(char letters: word){
            if(current->children[letters-'a'] == nullptr){
                current->children[letters-'a'] = new Trie();
            }
            current = current->children[letters-'a'];
        }
        current->isWord = true;
    }
    
    void search(Trie *root, string currentWord, vector<string> &result){
        if(result.size() == 3){
            return;
        }
        if(root->isWord){
            result.push_back(currentWord);
        }
        for(int i=0;i!=26;++i){
            if(root->children[i] != nullptr){
                string temp(1, i+'a');
                search(root->children[i],currentWord+temp,result);
            }
        }
    }
    
    vector<string> getSuggestions(string word){
        vector<string> result;
        Trie* current = this;
        for(char letters: word){
            if(current->children[letters-'a'] == nullptr){
                return result;
            }
            current = current->children[letters-'a'];
        }
        search(current, word, result);
        return result;
    }  
};
class Solution {
public:
    vector<vector<string>> suggestedProducts(vector<string>& products, string searchWord) {
        Trie* root = new Trie();
        for(string words: products){
            root->insert(words);
        }
        vector<vector<string>> result;
        string currentWord = "";
        for(char letters: searchWord){
            currentWord += letters;
            result.push_back(root->getSuggestions(currentWord));
        }
        return result;
    }
};